## Start up corda network


### 1 Run code nodes
from your cli folder, run following commands, check logs under network/corda/nodes/*/log to verify that each node has been started

> * ./network_clean.sh to clean up logs and vault database
> * ./network_start.sh to start up nodes

to stop all nodes, run following command
> * ./network_stop.sh to shutdown all nodes

### 2 Start up web clients

from your cli folder, run following command to start up all web servers, check cli/log to verify web servers are started.

> * ./start_webservers.sh

from your cli folder, run following command to stop all web servers
> * ./stop_webservers.sh

### 3 Open swagger UI

* alice: http://localhost:9000/swagger-ui.html
* charlie : http://localhost:9001/swagger-ui.html
* bob: http://localhost:9002/swagger-ui.html
* bank: http://localhost:9004/swagger-ui.html

