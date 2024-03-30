FROM centos:8
COPY bin/eproxy /eproxy/
COPY ebpf/output/connect.o ebpf/