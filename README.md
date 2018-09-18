### docker scripts and the runner

the runner will be installed inside each docker instance (that mean whenever having change in runner), we must
  - build runner on local
  - update ancestor dockerfile for copy runner binary into it
  - rebuild and push all dockerfile (ancestor must be the first)

### the project templates

the project templates will be differented depended on languages
  - templates
  - code implementation the abstract project generator (so that runner will call later)
 
We need it as simple as possible

