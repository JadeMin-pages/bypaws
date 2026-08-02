FUNCTION BUILD($CONFIG) {
	WRITE-HOST "Downloading dependencies..."
	go mod tidy
	WRITE-HOST "DONE"
	WRITE-HOST ""

	WRITE-HOST "Building project..."
	FOREACH ($OS in $CONFIG.BUILD.Keys) {
		FOREACH ($ARCH in $CONFIG.BUILD[$OS]) {
			$FILENAME = "$( $CONFIG.PROJECT_NAME )-$OS-$ARCH"
			IF ($OS -EQ "windows") {
				$FILENAME = "$FILENAME.exe"
			}
			
			WRITE-HOST "$FILENAME - " -NoNewline

			$ENV:GOOS = $OS
			$ENV:GOARCH = $ARCH
			go build -gcflags=all="-B" -ldflags="-s -w" -trimpath -o="bin/$FILENAME" ./src

			WRITE-HOST "DONE"
		}
	}
}



BUILD @{
	PROJECT_NAME = "bypaws"
	BUILD = @{
		"windows" = @("amd64")
		"linux" = @("amd64")
	}
}