### TODO:

1. Modules
- [ ] abstract project template/structure **-(WIP)-**
- [ ] abstract code compiler **-(WIP)-**
- [ ] abstract result handler (stdout/stderr) **-(WIP)-**

- [ ] code runner **-(WIP)-**
- [ ] entry cli
- [ ] entry http/2 server
- [ ] entry ipfs-pub-sub client/server for sharing peer-servers resources (https://github.com/ipfs/go-ipfs)

- [ ] docker scripts and the runner **-(WIP)-**
    
    the runner will be installed inside each docker instance (that mean whenever having change in runner), we must
  - ~build runner on local~
  - ~update ancestor dockerfile for copy runner binary into it~
  - rebuild and push all dockerfile (ancestor must be the first)

- [ ] the project templates **-(WIP)-**

    the project templates will be differented depended on languages
  - templates
  - code implementation the abstract project generator (so that runner will call later)

2. Development

  1. Build ancestor with coflies built by previous image of runner
  
```bash
$./build.sh
```

  2. Build coflies on local machine

```bash
$make all
```

NOTE: We need it as simple as possible

NOTE 02: When prototyping I put lot of thing together - We need to break them out into other folders/files when impl
