CREATE TRIGGER oi_chg_insert AFTER INSERT ON oi_chg FOR EACH ROW EXECUTE PROCEDURE on_oi_chg_insert();
CREATE TRIGGER volume_insert AFTER INSERT ON volume FOR EACH ROW EXECUTE PROCEDURE on_volume_insert();
CREATE TRIGGER oorders_insert AFTER INSERT ON oorders FOR EACH ROW EXECUTE PROCEDURE on_oorders_insert();
CREATE TRIGGER oorders_delete AFTER DELETE ON oorders FOR EACH ROW EXECUTE PROCEDURE on_oorders_delete();
CREATE TRIGGER mkt_stats_insert AFTER INSERT ON market FOR EACH ROW EXECUTE PROCEDURE on_market_insert();
CREATE TRIGGER mkt_stats_delete AFTER DELETE ON market FOR EACH ROW EXECUTE PROCEDURE on_market_delete();
CREATE TRIGGER mktord_insert AFTER INSERT ON mktord FOR EACH ROW EXECUTE PROCEDURE on_mktord_insert();
CREATE TRIGGER mkkord_delete AFTER DELETE ON mktord FOR EACH ROW EXECUTE PROCEDURE on_mktord_delete();
CREATE TRIGGER trd_mkt_stats_insert AFTER INSERT ON trd_mkt_stats FOR EACH ROW EXECUTE PROCEDURE on_trd_mkt_stats_insert();
CREATE TRIGGER trd_mkt_stats_delete AFTER DELETE ON trd_mkt_stats FOR EACH ROW EXECUTE PROCEDURE on_trd_mkt_stats_delete();
CREATE TRIGGER trd_mkt_stats_update AFTER UPDATE ON trd_mkt_stats FOR EACH ROW EXECUTE PROCEDURE on_trd_mkt_stats_update();
CREATE TRIGGER mktfin_insert AFTER INSERT on mkt_fin FOR EACH ROW EXECUTE PROCEDURE on_mktfin_insert();
CREATE TRIGGER mktfin_delete AFTER DELETE on mkt_fin FOR EACH ROW EXECUTE PROCEDURE on_mktfin_delete();
CREATE TRIGGER pl_insert AFTER INSERT on profit_loss FOR EACH ROW EXECUTE PROCEDURE on_profit_loss_insert();
CREATE TRIGGER pl_delete AFTER DELETE on profit_loss FOR EACH ROW EXECUTE PROCEDURE on_profit_loss_delete();
CREATE TRIGGER pl_update AFTER UPDATE on profit_loss FOR EACH ROW EXECUTE PROCEDURE on_profit_loss_update();
CREATE TRIGGER report_insert AFTER INSERT on report FOR EACH ROW EXECUTE PROCEDURE on_report_insert();
CREATE TRIGGER report_delete AFTER DELETE on report FOR EACH ROW EXECUTE PROCEDURE on_report_delete();
