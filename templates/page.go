package templates

var PageTpl = `
{{- define "page" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
{{ if gt (len .VennCharts) 0 }}
    <script type="text/javascript">
    "use strict";
    var tooltip = d3
        .select("body")
        .append("div")
        .attr("class", "venntooltip");

    function vennTooltip(div) {
        div.selectAll("g")
				.on("mouseover", function (d, i) {
						/* sort all the areas relative to the current item */ 
						venn.sortAreas(
                                div,
								d
						);
						/* Display a tooltip with the current size */ tooltip
								.transition()
								.duration(400)
								.style("opacity", 0.9);
						tooltip.text(d.size + " users");
						/* highlight the current path */ var selection = d3
								.select(this)
								.transition("tooltip")
								.duration(400);
						selection
								.select("path")
								.style("fill-opacity", d.sets.length == 1 ? 0.4 : 0.1)
								.style("stroke-opacity", 1);
				})
				.on("mousemove", function () {
						tooltip
								.style("left", d3.event.pageX + "px")
								.style("top", d3.event.pageY - 28 + "px");
				})
				.on("mouseout", function (d, i) {
						tooltip.transition().duration(400).style("opacity", 0);
						var selection = d3
								.select(this)
								.transition("tooltip")
								.duration(400);
						selection
								.select("path")
								.style("fill-opacity", d.sets.length == 1 ? 0.25 : 0.0)
								.style("stroke-opacity", 0);
				});
    }
    </script>
{{ end }}

{{ if eq .Layout "none" }}
    {{- range .Charts }} {{ template "base" . }} {{- end }}
    {{- range .VennCharts }} {{ template "venn" . }} {{- end }}
{{ end }}

{{ if eq .Layout "center" }}
    <style> .container {display: flex;justify-content: center;align-items: center;} .item {margin: auto;} </style>
    {{- range .Charts }} {{ template "base" . }} {{- end }}
    {{- range .VennCharts }} {{ template "venn" . }} {{- end }}
{{ end }}

{{ if eq .Layout "flex" }}
    <style> .box { justify-content:center; display:flex; flex-wrap:wrap } </style>
    <div class="box"> {{- range .Charts }} {{ template "base" . }} {{- end }} </div>
    <div class="box"> {{- range .VennCharts }} {{ template "venn" . }} {{- end }} </div>
{{ end }}
</body>
</html>
{{ end }}
`
