{ {{ range getEnvArray "KEYS" }}
	{
		"key": "{{ . }}",
		"value": "{{ getEnvMap "MAP" . }}",
	},{{ end }}
}
