echo "Building ./gopher-ball (debug version)"

go build -ldflags "-s" -o gopher-ball

if [ $? -gt 0 ]
then
  echo "Basic build failed"
  exit 1
fi

echo "./gopher-ball created (ignored by git)"

echo "Building ./gopher-ball.app"

go build -ldflags "-s" -o gopher-ball.app

if [ $? -gt 0 ]
then
  echo "Mac build failed"
  exit 1
fi

echo "./gopher-ball.app created (ignored by git)"

exit 0