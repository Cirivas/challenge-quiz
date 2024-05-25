# Variable for filename for store running procees id
PID_FILE = /tmp/backend-api.pid

# Start task performs build command and writes it's process id to PID_FILE.
start:
	go build -o /bin/backend-api . && \
	/bin/backend-api & echo $$! > $(PID_FILE)

# Stop task will kill process by ID stored in PID_FILE (and all child processes by pstree).  
stop:
	@echo "KILL backend-api" && printf '%*s\n' "40" '' | tr ' ' -
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 
  
# Restart task will execute stop and start tasks in strict order and prints message. 
restart: stop start
	@echo "STARTED backend-api" && printf '%*s\n' "40" '' | tr ' ' -
  
# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
serve: start
	fswatch -m poll_monitor -xor --event Updated --event Created . | xargs -n1 -I {} make restart
#	xargs -n1 -I{} make restart
  
# .PHONY is used for reserving tasks words
.PHONY: start stop restart serve