trap 'kill 0' EXIT
export PRODUCTION=false
cd personal-budget-app-backend
echo "Running backend in development mode"
go run . &
cd ..
cd personal-budget-app-frontend
npm run dev &
cd ..
wait
