// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package tables

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"app/internal/db"
	"fmt"
	"github.com/jackc/pgx/pgtype"
)

func MatchdayPredictions(matches []db.Match, predictions []db.UserPredictions) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"overflow-x-auto w-full\"><div class=\"inline-block min-w-full rounded-lg overflow-hidden shadow-lg\"><table class=\"min-w-full bg-gray-800 text-sm border-collapse\"><thead class=\"bg-gray-800\"><tr><th class=\"sticky left-0 z-10 bg-gray-800 py-3 px-4 border-b-2 border-r-2 border-gray-600 rounded-tl-lg\">Partidos</th>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, pred := range predictions {
			if i == len(predictions)-1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<th class=\"py-3 px-1 border-b-2 border-gray-600 text-center rounded-tr-lg\" colspan=\"2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(pred.UserName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 19, Col: 50}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<th class=\"py-3 px-1 border-b-2 border-r-2 border-gray-600 text-center\" colspan=\"2\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(pred.UserName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 23, Col: 50}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, match := range matches {
			var templ_7745c5c3_Var4 = []any{templ.KV("bg-gray-750", i%2 == 0), templ.KV("bg-gray-800", i%2 != 0), "hover:bg-gray-800"}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var4).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><td class=\"sticky left-0 z-10 py-2 px-2 border-b border-r border-gray-600\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%s vs %s", match.HomeName, match.AwayName))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 33, Col: 88}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for j, pred := range predictions {
				if i < len(pred.Predictions) {
					if i == len(matches)-1 {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var7 = []any{
							"py-2 px-2 border-b border-r border-gray-600 text-center",
							templ.KV("bg-green-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 0 && pred.IsCorrect[i][0].Status == pgtype.Present && pred.IsCorrect[i][0].Bool),
							templ.KV("bg-red-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 0 && pred.IsCorrect[i][0].Status == pgtype.Present && !pred.IsCorrect[i][0].Bool),
						}
						templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var7...)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var8 string
						templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var7).String())
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 1, Col: 0}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" colspan=\"2\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var9 string
						templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(pred.Predictions[i][0])
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 44, Col: 67}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var10 = []any{
							"py-2 px-2 border-b border-gray-600 text-center",
							templ.KV("bg-green-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 0 && pred.IsCorrect[i][0].Status == pgtype.Present && pred.IsCorrect[i][0].Bool),
							templ.KV("bg-red-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 0 && pred.IsCorrect[i][0].Status == pgtype.Present && !pred.IsCorrect[i][0].Bool),
						}
						templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var10...)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var11 string
						templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var10).String())
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 1, Col: 0}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var12 string
						templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(pred.Predictions[i][0])
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 53, Col: 67}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if j == len(predictions)-1 {
							var templ_7745c5c3_Var13 = []any{
								"py-2 px-2 border-b border-gray-600 text-center",
								templ.KV("bg-green-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 1 && pred.IsCorrect[i][1].Status == pgtype.Present && pred.IsCorrect[i][1].Bool),
								templ.KV("bg-red-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 1 && pred.IsCorrect[i][1].Status == pgtype.Present && !pred.IsCorrect[i][1].Bool),
							}
							templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var13...)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var14 string
							templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var13).String())
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 1, Col: 0}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var15 string
							templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(pred.Predictions[i][1])
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 61, Col: 71}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						} else {
							var templ_7745c5c3_Var16 = []any{
								"py-2 px-2 border-b border-r border-gray-600 text-center",
								templ.KV("bg-green-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 1 && pred.IsCorrect[i][1].Status == pgtype.Present && pred.IsCorrect[i][1].Bool),
								templ.KV("bg-red-800", len(pred.IsCorrect) > i && len(pred.IsCorrect[i]) > 1 && pred.IsCorrect[i][1].Status == pgtype.Present && !pred.IsCorrect[i][1].Bool),
							}
							templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var16...)
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var17 string
							templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var16).String())
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 1, Col: 0}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var18 string
							templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(pred.Predictions[i][1])
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/components/shared/tables/matchday_predictions.templ`, Line: 69, Col: 71}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
					}
				} else {
					if j == len(predictions)-1 {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"py-2 px-2 border-b border-gray-600 text-center\" colspan=\"2\"></td>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"py-2 px-2 border-b border-r-2 border-gray-600 text-center\" colspan=\"2\"></td>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
