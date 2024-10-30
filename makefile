# 실행 파일 이름 설정
BINARY_NAME=gogo.exe
GO_FILES=main.go weather.go

# 기본 명령: Windows용 exe 파일 빌드
build:
	go build -o $(BINARY_NAME) $(GO_FILES)

# 정리: 빌드된 파일 삭제
clean:
	del /f $(BINARY_NAME)
