FROM golang:1.24 as builder

COPY . /teacup
WORKDIR /teacup/mayhem

RUN GOTOOLCHAIN=auto go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOTOOLCHAIN=auto go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN GOTOOLCHAIN=auto go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_switchboard_event && go-fuzz-build -libfuzzer -o fuzz_switchboard_event.a && \
    clang -fsanitize=fuzzer fuzz_switchboard_event.a -o fuzz_switchboard_event.libfuzzer

RUN cd fuzz_switchboard_utils && go-fuzz-build -libfuzzer -o fuzz_switchboard_utils.a && \
    clang -fsanitize=fuzzer fuzz_switchboard_utils.a -o fuzz_switchboard_utils.libfuzzer

RUN cd fuzz_switchboard_watcher && go-fuzz-build -libfuzzer -o fuzz_switchboard_watcher.a && \
    clang -fsanitize=fuzzer fuzz_switchboard_watcher.a -o fuzz_switchboard_watcher.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /teacup/mayhem/fuzz_switchboard_event/fuzz_switchboard_event.libfuzzer /
COPY --from=builder /teacup/mayhem/fuzz_switchboard_utils/fuzz_switchboard_utils.libfuzzer /
COPY --from=builder /teacup/mayhem/fuzz_switchboard_watcher/fuzz_switchboard_watcher.libfuzzer /

