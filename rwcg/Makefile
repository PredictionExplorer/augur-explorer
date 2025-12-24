.PHONY: all clean etl websrv cg_etl rw_etl notibot tools rw_tools

all: etl websrv notibot

etl: cg_etl rw_etl

cg_etl:
	cd etl/cosmicgame && go build -o cg_etl .

rw_etl:
	cd etl/randomwalk && go build -o rw_etl .

websrv:
	cd websrv && go build -o websrv .

notibot:
	cd notibot && go build -o notibot .

tools:
	cd tools && go build -o evtlog_diff evtlog_diff.go
	cd tools && go build -o db_verify db_verify.go
	cd tools && go build -o archive_export archive_export.go

rw_tools:
	cd etl/randomwalk/tools && go build -o rw_toprated rw_toprated.go

clean:
	rm -f etl/cosmicgame/cg_etl
	rm -f etl/randomwalk/rw_etl
	rm -f etl/randomwalk/tools/rw_toprated
	rm -f websrv/websrv
	rm -f notibot/notibot
	rm -f notibot/scripts/verify-token-imgs
	rm -f tools/evtlog_diff
	rm -f tools/db_verify
	rm -f tools/archive_export
