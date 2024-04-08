//go:build ignore
#include "vmlinux.h"
#include "socket.h"
#include "bpf_helpers.h"
#include "bpf_endian.h"

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("cgroup/connect4")
int connect4(struct bpf_sock_addr *ctx) {
    int ret = 1; /* OK value */
    if (ctx->type != SOCK_STREAM && ctx->type != SOCK_DGRAM) {
        bpf_printk("unkonw socket type");
        return ret;
    }

    __u8 ip_proto;
    switch (ctx->type) {
    case SOCK_STREAM:
        ip_proto = IPPROTO_TCP;
        break;
    case SOCK_DGRAM:
        ip_proto = IPPROTO_UDP;
        break;
    default:
        return ret;
    }
    return 1;
}
