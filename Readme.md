# binarydist

Package binarydist implements binary diff and patch as described on
<http://www.daemonology.net/bsdiff/>. It reads and writes files
compatible with the tools there.

This project is forked from <https://github.com/kr/binarydist>
The reason why we have to fork this proeject is that, in kr/binarydist, the authors had to fork a process to do the bzip compression, because the native golang lib only have uncompressor, not compressor.
Doing such fork is ok in general usage. However, in a such heavy loaded system, forking a process on the fly will worse hardware resources, which reduces the overall throughput.
Instead of doing that, we decided to use cgo to wrap bzip compressor which could be used to replace the forked one, the performance should be better.

If you have any bugs, please mail to yhh92u@gmail.com
