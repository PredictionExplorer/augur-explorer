## ETL (Extract Transform Load) daemon for converting blockchain events to SQL database

#### Features

1. No-loss (reliable, failure proof) event gathering process. 
2. Only transactions with events are stored, transactions without events are dropped
3. Chain reorgs (with block deletes) are suported with SQL triggers
4. Acceleration via custom GetBlockReceipts call to Geth (in located in `./geth-patch` directory) available on demand


#### Basic functioning

The daemon processes blocks sequentially one by one and inserts transactions in `transaction` table and events in `evt_log` table. Block are validated using parent block hash so all the events are inserted in order. In case network failure/software failure occurs the daemon will delete incomplete state and reporcess partially inserted block. All tables dependent on `evt_log` must reference this table via 'REFERENCES' SQL keyword so upon (block) record DELETE (due to chain reorg) your data will be deleted and reinserted again (if aplicable). All deletes happen backwards (latest are deleted first).
Layer1 contains only transactions and events, your application should be built upon this layer as Layer 2


#### Installation

##### Create a Unix account to run the daemon using this account (i.e. 'myapp'):

    useradd -m 'myapp'
    (setup the password and configure the shell)

##### Create postgres user for this new account:

    su - postgres
    psql
    postgres-#  CREATE ROLE myapp WITH LOGIN CREATEDB ENCRYPTED PASSWORD 'myapp_pass';
    \q

    su - myapp
    createdb myappdb

##### Execute the following SQL scripts:

    cd ../../sql/layer1/
    psql < tables.sql
    psql < indices.sql
    psql < triggers-layer1.sql

##### Create configuration directory:

    mkdir ~/config

In this directory, create file 'etl-config.env' with the following content:

        export RPC_URL="http://127.0.0.1:8545"
        export EXTRACTOR_HOST="127.0.0.1:5432"
        export EXTRACTOR_USERNAME="myapp"
        export EXTRACTOR_DATABASE="myappdb"
        export EXTRACTOR_PASSWORD="myapp_pass"

##### Build executables (in cmd/layer1 directory):

	cd cmd/layer1
    go build .

##### Load environment variables (the file created in the step above):

    . ~/config/etl-config.env

##### Run Layer1 ETL daemon

    ./layer1

The log files will be created in $HOME/ae_logs:

    etl_info.log
    etl_error.log
    etl_db.log


Layer1 will populate tables `block`,`transaction`,`evt_log` , and your App should create new tables to process events with signatures of your choice. Your app daemon should be polling events from the DB and processing them on its own (separately from layer1).
