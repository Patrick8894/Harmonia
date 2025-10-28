# Hello
python clients/hello_client.py --name "Patrick"

# Evaluate
python clients/evaluate_client.py --expr "2 + 3*x - sqrt(y)" --vars '{"x":4,"y":9}'

# Transform (MAP)
python clients/transform_client.py --op MAP --expr "2*x+1" --data "[1,2,3,4]"

# Transform (FILTER)
python clients/transform_client.py --op FILTER --expr "x % 2 == 0" --data "[1,2,3,4,5,6]"

# Transform (SUM)
python clients/transform_client.py --op SUM --expr "x*x" --data "[1,2,3]"

# PlanTasks
python clients/plan_client.py --goal "Add Transform RPC and Dockerize" --hint tests --hint logging --max_steps 7
