### TODO:

1. Modules
- [ ] abstract project template/structure
- [ ] abstract code compiler
- [ ] abstract result handler (stdout/stderr)

- [ ] code runner
- [ ] entry cli
- [ ] entry http/2 server

- [ ] docker scripts and the runner **-(WIP)-**
    
    the runner will be installed inside each docker instance (that mean whenever having change in runner), we must
  - build runner on local
  - update ancestor dockerfile for copy runner binary into it
  - rebuild and push all dockerfile (ancestor must be the first)

- [ ] the project templates **-(WIP)-**

    the project templates will be differented depended on languages
  - templates
  - code implementation the abstract project generator (so that runner will call later)

2. Development

  1. Build Docker images
  
```bash
$./build.sh
```

  2. Build coflies on local machine

```bash
$make all
```

NOTE: We need it as simple as possible

NOTE 02: When prototyping I put lot of thing together - We need to break them out into other folders/files when impl
