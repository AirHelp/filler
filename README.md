# filler

Filler is a simple tool to fill config templates using environment variables. 


## CLI usage
### Usage:
```
filler --dir <directory> --ext <template_extension>
```

### Sample usage:
```
filler --dir /etc --ext tpl
```

sample test.conf.tpl placed in /etc :
```
This is template file version {{ getEnv "VERSION" }}
```

1. This sample command runs filler on /etc directory
2. Filler will search for files with mask *.devops_template (recursively)
3. Filler will find test.conf.devops_template
4. Fills every action getEnv with value of defined variable name.
5. Saves in file without template extension (in this case test.conf)
6. Removes template files
