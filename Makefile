all:
	@echo "  CC\tBridge"
	@gocross arm go build Bridge.go 2>/dev/null
	@echo "  CC\tAccessory"
	@gocross arm go build Accessory.go 2>/dev/null
	@echo "  STRIP\tBridge"
	@arm-linux-gnueabihf-strip Bridge
	@echo "  STRIP\tAccessory"
	@arm-linux-gnueabihf-strip Accessory
install: all
	@echo "  INST\tBridge"
	@install -m 0755 Bridge ../orangepi_zero/tools/Bridge
	@echo "  INST\tAccessory"
	@install -m 0755 Accessory ../orangepi_zero/tools/Accessory
clean:
	@echo "  CL\tBridge"
	@rm Bridge
	@echo "  CL\tAccessory"
	@rm Accessory
