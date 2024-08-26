// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "distro-hub/views/auth"

func Auth() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n\t\t\t.selected {\n\t\t\t\tbackground-color: #d5d5d5;\n\t\t\t}\n\t\t</style> <section class=\"min-h-screen grid place-items-center\"><article class=\"flex flex-col gap-5\"><div class=\"flex gap-2 border rounded border-gray-300 bg-gray-100 p-1\" id=\"tabs\" role=\"tablist\"><button hx-get=\"/auth/login\" hx-target=\"#form\" class=\"text-sm py-1 flex-1 hover:bg-gray-300 transition rounded selected\" role=\"tab\" aria-controls=\"login-tab\" aria-selected=\"true\">Log in</button> <button hx-get=\"/auth/register\" hx-target=\"#form\" class=\"text-sm py-1 flex-1 hover:bg-gray-300 transition rounded\" role=\"tab\" aria-controls=\"register-tab\" aria-selected=\"false\">Register</button></div><hr><div id=\"form\" class=\"\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = auth.LoginTab().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></article></section><script>\n\t\t\tdocument.addEventListener(\"DOMContentLoaded\", () => {\n\t\t\t\tlet currentTab;\n\t\t\t\tconst tabs = document.querySelectorAll(\"[role=tab]\")\n\t\t\t\ttabs.forEach(tab => {\n\t\t\t\t\ttab.addEventListener(\"click\", (event) => {\n\t\t\t\t\t\tcurrentTab = document.querySelector(\"[aria-selected=true]\")\n\t\t\t\t\t\tcurrentTab.setAttribute(\"aria-selected\", \"false\")\n\t\t\t\t\t\tcurrentTab.classList.remove(\"selected\")\n\t\t\t\t\t\tconst newTab = event.target\n\t\t\t\t\t\tnewTab.setAttribute(\"aria-selected\", \"true\")\n\t\t\t\t\t\tnewTab.classList.add(\"selected\")\n\t\t\t\t\t})\n\t\t\t\t})\n\t\t\t})\n\t\t</script> <noscript><p>Please enable JavaScript to use the login and registration tabs.</p></noscript>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base("Distro Hub | Auth").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
