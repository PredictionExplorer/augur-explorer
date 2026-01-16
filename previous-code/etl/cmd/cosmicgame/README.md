## ETL (Extract Transform Load) daemon for Cosmic Signature Token

#### Dependencies:

	Layer 1 ETL daemon located in adjacent directory (../cmd/layer1)
	
#### Installation

##### Create a Unix account to run the daemon (i.e. 'cosmicgame'):

    useradd -m 'cosmicgame'
    (setup the password and configure the shell)

##### Create postgres user for this new account:

    su - postgres
    psql
    postgres-#  CREATE ROLE cosmicgame WITH LOGIN CREATEDB ENCRYPTED PASSWORD 'cosmicgame_pass';
    \q

    su - cosmicgame
    createdb cosmicgamedb

##### Execute the following SQL scripts:

    cd ../../sql/cosmicgame/
    psql < tables-cosmicgame.sql
    psql < indices_cosmicgame.sql
    psql < cosmicgame-funcs.sql
    psql < triggers-cosmicgame.sql

(the tables of Layer 1 ETL should be already there, if not, follow instructions at ../cmd/layer1)

##### Create configuration directory:

    mkdir ~/config

In this directory, create file 'etl-config.env' with the following content:

        export RPC_URL="http://127.0.0.1:8545"
        export EXTRACTOR_HOST="127.0.0.1:5432"
        export EXTRACTOR_USERNAME="cosmicgame"
        export EXTRACTOR_DATABASE="cosmicgamedb"
        export EXTRACTOR_PASSWORD="cosmicgame_pass"

(Replace values of the variables with values of your configuration)

##### Build executables (in cmd/cosmicgame directory):

	cd cmd/cosmicgame
    go build .

##### Load environment variables (the file created in the step above):

    . ~/config/etl-config.env

##### Run CosmicGame ETL daemon

    ./cosmicgame

The log files will be created in $HOME/ae_logs:

    cosmicgame_info.log
    cosmicgame_error.log
    cosmicgame_db.log


