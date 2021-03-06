[#ftl/]
[#-- @ftlvariable name="activationComplete" type="boolean" --]
[#-- @ftlvariable name="application" type="io.fusionauth.domain.Application" --]
[#-- @ftlvariable name="client_id" type="java.lang.String" --]
[#-- @ftlvariable name="devicePendingIdPLink" type="io.fusionauth.domain.provider.PendingIdPLink" --]
[#-- @ftlvariable name="tenant" type="io.fusionauth.domain.Tenant" --]
[#-- @ftlvariable name="tenantId" type="java.util.UUID" --]
[#-- @ftlvariable name="userCodeLength" type="int" --]
[#-- @ftlvariable name="version" type="java.lang.String" --]
[#import "../_helpers.ftl" as helpers/]

[@helpers.html]
  [@helpers.head title=theme.message('device-title')]
  <script src="/js/oauth2/Device.js?version=${version}"></script>
  <script>
    Prime.Document.onReady(function() {
      var form = Prime.Document.queryById('device-form');
      new FusionAuth.OAuth2.Device(form, ${userCodeLength});
    });
  </script>
  <style>
    #user_code_container {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
    }

    #user_code_container > div {
      margin-left: 2px;
      margin-right: 2px;
    }

    #user_code_container input[type="text"] {
      font-size: 30px;
      padding: 5px 0;
      margin-bottom: 5px;
      text-align: center;
      width: 32px;
    }

    #user_code_container input[type="text"] + span {
      font-size: 32px;
    }
  </style>
  [/@helpers.head]
  [@helpers.body]

    [@helpers.header]
      [#-- Custom header code goes here --]
    [/@helpers.header]

    [@helpers.main title=theme.message('device-form-title')]
      [#setting url_escaping_charset='UTF-8']
      [#-- During a linking work flow, optionally indicate to the user which IdP is being linked. --]
      [#if devicePendingIdPLink??]
        <p class="mt-0">
          ${theme.message('pending-device-link', devicePendingIdPLink.identityProviderName)}
        </p>
      [/#if]
      <form action="/oauth2/device" method="POST" id="device-form">
        [@helpers.oauthHiddenFields/]
        <p>${theme.message('userCode')}</p>
        <fieldset>
          <div id="user_code_container">
            [#list 0..<userCodeLength as i]
              <div>
                <label for="user_code_${i}"></label>
                <input type="text" id="user_code_${i}" maxlength="1" [#if i?index == 0]autofocus[/#if] autocomplete="off"/>
                [#if i == (userCodeLength/2)?floor - 1]<span>-</span>[/#if]
              </div>
            [/#list]
            <input type="hidden" name="interactive_user_code" id="interactive_user_code" />
          </div>
        </fieldset>

        <div class="form-row">
          [@helpers.errors field="user_code" /]
        </div>

        <div class="form-row push-top">
          [@helpers.button text=theme.message('submit')/]
        </div>
      </form>
    [/@helpers.main]

    [@helpers.footer]
      [#-- Custom footer code goes here --]
    [/@helpers.footer]
  [/@helpers.body]
[/@helpers.html]