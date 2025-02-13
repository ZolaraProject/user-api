{{- define "userapi.name" -}}
{{- default "user-api" .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Helm required labels */}}
{{- define "userapi.labels" -}}
heritage: {{ .Release.Service }}
release: {{ .Release.Name }}
chart: {{ .Chart.Name }}
app: "{{ template "userapi.name" . }}"
{{- end -}}

{{/* matchLabels */}}
{{- define "userapi.matchLabels" -}}
release: {{ .Release.Name }}
app: "{{ template "userapi.name" . }}"
{{- end -}}