package templates

var VennTpl = `
{{- define "venn" }}
<div id="venn_{{ .ChartID | safeJS }}"></div>

<script type="text/javascript">
    "use strict";
		var sets_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};

		var chart_{{ .ChartID | safeJS }} = venn.VennDiagram().width({{ .Width | safeJS }}).height({{ .Heigth | safeJS }});
		var div_{{ .ChartID | safeJS }} = d3.select("#venn_{{ .ChartID | safeJS }}");
		div_{{ .ChartID | safeJS }}.datum(sets_{{ .ChartID | safeJS }}).call(chart_{{ .ChartID | safeJS }});
		div_{{ .ChartID | safeJS }}.selectAll("path")
				.style("stroke-opacity", 0)
				.style("stroke", "#fff")
				.style("stroke-width", 3);
		vennTooltip(div_{{ .ChartID | safeJS }});
</script>
{{ end }}
`
