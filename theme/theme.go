package theme

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
)

const defaultMessagesSeperator = "\n"

func LoadFromDisk(path string, defaultMessageSubstitutions map[string]string) (*fusionauth.Theme, error) {
	theme := &fusionauth.Theme{}
	templates := &fusionauth.Templates{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fname := strings.Split(f.Name(), ".")
		switch fname[1] {
		case "ftl":
			// template
			v := reflect.ValueOf(templates).Elem().FieldByName(fname[0])
			if v.IsValid() {
				data, err := os.ReadFile(filepath.Join(path, f.Name()))
				if err != nil {
					return nil, err
				}
				v.SetString(string(data))
			}
		case "css":
			// stylesheet
			data, err := os.ReadFile(filepath.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}
			theme.Stylesheet = string(data)
		case "conf":
			// default messages
			data, err := os.ReadFile(filepath.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}
			theme.DefaultMessages = insertSubstitutions(string(data), defaultMessageSubstitutions)
		default:
			log.Printf("[ERROR] unexpected file found %v", f.Name())
		}
	}
	theme.Templates = *templates

	return theme, nil
}

func WriteToDisk(theme *fusionauth.Theme, path string, defaultMessageSubstitutions map[string]string) error {
	os.Mkdir(path, os.ModePerm)

	// Write out the templates
	v := reflect.ValueOf(theme.Templates)
	typeOfTheme := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fname := filepath.Join(path, typeOfTheme.Field(i).Name)

		f, err := os.Create(fmt.Sprintf("%v.ftl", fname))
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(v.Field(i).String())
	}

	// Write out the stylesheet
	f, err := os.Create(filepath.Join(path, "stylesheet.css"))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(theme.Stylesheet)

	// Write out the default messages
	f, err = os.Create(filepath.Join(path, "DefaultMessages.conf"))
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(removeSubstitutions(theme.DefaultMessages, defaultMessageSubstitutions))

	return nil
}

func LoadSubstitutionsFromDisk(path string) (map[string]string, error) {
	substitutions := make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		return substitutions, nil
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), "=")
		substitutions[l[0]] = l[1]
	}

	return substitutions, nil
}

func insertSubstitutions(raw string, subs map[string]string) string {
	defaultMessages := strings.Split(raw, defaultMessagesSeperator)
Substitutions:
	for subK, subV := range subs {
		log.Printf("[INFO] subbing %v=%v", subK, subV)
		// Search defaultMessages for the substitution
		for i, v := range defaultMessages {
			log.Printf("[DEBUG] trying %v", v)
			blankSub := fmt.Sprintf("%v=", subK)
			if strings.HasPrefix(v, blankSub) {
				defaultMessages[i] = fmt.Sprintf("%v=%v", subK, subV)
				continue Substitutions
			}
		}
		log.Printf("[ERROR] could not substitute %v", subK)
	}

	return strings.Join(defaultMessages, defaultMessagesSeperator)
}

func removeSubstitutions(raw string, subs map[string]string) string {
	defaultMessages := strings.Split(raw, defaultMessagesSeperator)
Substitutions:
	for k, _ := range subs {
		// Search defaultMessages for the substitution
		for i, v := range defaultMessages {
			blankSub := fmt.Sprintf("%v=", k)
			if strings.HasPrefix(v, blankSub) {
				defaultMessages[i] = blankSub
				continue Substitutions
			}
		}
		log.Printf("[ERROR] could not substitute %v", k)
	}

	return strings.Join(defaultMessages, defaultMessagesSeperator)
}
