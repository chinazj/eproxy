//go:build ignore
#include "vmlinux.h"
#include "socket.h"
#include "bpf_helpers.h"
#include "bpf_endian.h"

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("cgroup/connect4")
int connect4(struct bpf_sock_addr *ctx) {
    struct sockaddr_in sa = {};
    struct svc_addr *orig;

    /* Force local address to 127.0.0.1:22222. */
    sa.sin_family = AF_INET;
    sa.sin_port = bpf_htons(22222);
    sa.sin_addr.s_addr = bpf_htonl(0x7f000001);
    bpf_printk("connector 4\n");
//    if (bpf_bind(ctx, (struct sockaddr *)&sa, sizeof(sa)) != 0)
//        return 0;

//    /* Rewire service 1.2.3.4:60000 to backend 127.0.0.1:60123. */
    if (ctx->user_port == bpf_htons(60000)) {
        bpf_printk("user_port 6000\n");
//        orig = bpf_sk_storage_get(&service_mapping, ctx->sk, 0,
//                      BPF_SK_STORAGE_GET_F_CREATE);
//        if (!orig)
//            return 0;
//
//        orig->addr = ctx->user_ip4;
//        orig->port = ctx->user_port;
//
//        ctx->user_ip4 = bpf_htonl(0x7f000001);
//        ctx->user_port = bpf_htons(60123);
    }
    return 1;
}
