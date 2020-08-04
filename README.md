
# Dynamic DNS Updater for GratisDNS

This is a small tool to update DNS records managed by the danish provider
GratisDNS. https://web.gratisdns.dk/

You can find information aboutthe DDNS setup at https://larsendata.wiki/gratisdns:ddns

## Configuring this tool

The tool will read a file named `config.yaml` in the current working directory,
after updating all domains, the tool exits.

The configuration file should look like this:

```
username: YOUR_USER
password: THE_PASSWORD
domain_account: THE_ACCOUNT_FOR_THE_RECORDS
domains:
- foo.example.com
- bar.example.com
- etc
```
