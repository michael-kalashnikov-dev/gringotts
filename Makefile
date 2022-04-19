proto_gen:
	@scripts/proto/generate.sh

.PHONY: proto_gen

proto_clear:
	@scripts/proto/clear.sh

.PHONY: proto_clear