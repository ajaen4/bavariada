// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package graphs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func GraphContainer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full bg-white rounded-lg shadow dark:bg-gray-800 p-4 md:p-6\"><div class=\"flex justify-between mb-5\"><div class=\"grid gap-4 grid-cols-2\"><div><h5 class=\"inline-flex items-center text-gray-500 dark:text-gray-400 leading-none font-normal mb-2\">Jornadas jugadas</h5><p id=\"matchday-count\" class=\"text-gray-900 dark:text-white text-2xl leading-none font-bold\">Loading...</p></div></div></div><div id=\"chart-content\"><div id=\"line-chart\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = GraphScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func GraphScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        document.addEventListener('DOMContentLoaded', function() {\n            fetch('/points/cumulative')\n                .then(response => response.json())\n                .then(chartData => {\n                    document.getElementById('matchday-count').textContent = chartData.total_matchdays;\n\n                    const series = chartData.series.map(user => ({\n                        name: user.user_name,\n                        data: user.cumulative_points,\n                        color: user.color\n                    }));\n\n                    const options = {\n                        chart: {\n                            height: \"100%\",\n                            maxWidth: \"100%\",\n                            type: \"line\",\n                            fontFamily: \"Inter, sans-serif\",\n                            dropShadow: {\n                                enabled: false,\n                            },\n                            toolbar: {\n                                show: false,\n                            },\n                        },\n                        tooltip: {\n                            enabled: true,\n                            x: {\n                                show: false,\n                            },\n                        },\n                        dataLabels: {\n                            enabled: false,\n                        },\n                        stroke: {\n                            width: 3,\n                            curve: 'straight'\n                        },\n                        grid: {\n                            show: true,\n                            strokeDashArray: 4,\n                            padding: {\n                                left: 10,\n                                right: 10,\n                                top: -20,\n                                bottom: 20,\n                            },\n                        },\n                        series: series,\n                        colors: series.map(s => s.color),\n                        legend: {\n                            show: true,\n                            position: 'bottom',\n                            horizontalAlign: 'center',\n                            floating: true,\n                            offsetY: 10,\n                            offsetX: -5\n                        },\n                        xaxis: {\n                            categories: chartData.categories,\n                            labels: {\n                                show: true,\n                                style: {\n                                    fontFamily: \"Inter, sans-serif\",\n                                    cssClass: 'text-xs font-normal fill-gray-500 dark:fill-gray-400'\n                                }\n                            },\n                            axisBorder: {\n                                show: false,\n                            },\n                            axisTicks: {\n                                show: false,\n                            },\n                        },\n                        yaxis: {\n                            show: true,\n                            labels: {\n                                show: true,\n                                style: {\n                                    fontFamily: \"Inter, sans-serif\",\n                                    cssClass: 'text-xs font-normal fill-gray-500 dark:fill-gray-400'\n                                }\n                            },\n                        },\n                    };\n\n                    new ApexCharts(document.getElementById(\"line-chart\"), options).render();\n                })\n                .catch(error => console.error('Error fetching chart data:', error));\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
