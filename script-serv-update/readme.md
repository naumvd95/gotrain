Design the workflow and write the script to update hardware firmware on 30k servers,
3 DCs, 20 servers per rack, 1 rack can be offline for patching.


note:
for yaml struct use:      "gopkg.in/yaml.v2"

for routine error handling use
- channel hint https://golangcode.com/errors-in-waitgroups/
- or manage map with errors
