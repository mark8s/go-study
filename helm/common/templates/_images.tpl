{{/* vim: set filetype=mustache: */}}
{{/*
Return the proper image name
{{ include "common.images.image" ( dict "imageRoot" .Values.path.to.the.image "global" $) }}
*/}}
{{- define "common.images.image" -}}
{{- if  .Values.image }}
   {{- if .Values.image.name }}
      {{- .Values.image.name }}
   {{- else }}
      {{- .global }}
   {{- end }}
{{- else }}
  {{- .global }}
{{- end }}
{{- end }}

