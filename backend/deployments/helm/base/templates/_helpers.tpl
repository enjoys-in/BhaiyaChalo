{{- define "base.labels" -}}
app.kubernetes.io/name: {{ .Values.service.name }}
app.kubernetes.io/version: {{ .Values.image.tag | default "latest" }}
app.kubernetes.io/managed-by: helm
app.kubernetes.io/part-of: bhaiyachalo
{{- end }}

{{- define "base.selectorLabels" -}}
app.kubernetes.io/name: {{ .Values.service.name }}
{{- end }}

{{- define "base.image" -}}
{{ .Values.image.repository }}/{{ .Values.service.name }}:{{ .Values.image.tag | default "latest" }}
{{- end }}
