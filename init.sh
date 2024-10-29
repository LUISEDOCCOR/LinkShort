echo "LINK SHORTENER"

case $1 in
 --dev)
    go install github.com/air-verse/air@latest
    go mod tidy
    export PATH=$PATH:$HOME/go/bin
    clear
    air
 ;;
 --prod)
    go mod tidy
    clear
    go build -o ./build/main
    ./build/main
 ;;
esac
