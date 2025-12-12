.PHONY: all clean etl websrv cg_etl rw_etl

all: etl websrv

etl: cg_etl rw_etl

cg_etl:
	cd etl/cosmicgame && go build -o cg_etl .

rw_etl:
	cd etl/randomwalk && go build -o rw_etl .

websrv:
	cd websrv && go build -o websrv .

clean:
	rm -f etl/cosmicgame/cg_etl
	rm -f etl/randomwalk/rw_etl
	rm -f websrv/websrv
