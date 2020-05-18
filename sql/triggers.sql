CREATE TRIGGER update_oi_insert AFTER INSERT ON oi_chg FOR EACH ROW EXECUTE PROCEDURE update_oi_on_insert();
CREATE TRIGGER update_vol_insert AFTER INSERT ON volume FOR EACH ROW EXECUTE PROCEDURE update_vol_on_insert();
CREATE TRIGGER update_mkt_fin_insert AFTER INSERT ON mkt_fin FOR EACH ROW EXECUTE PROCEDURE update_mkt_fin_on_insert();
CREATE TRIGGER update_oostats_insert AFTER INSERT ON oorders FOR EACH ROW EXECUTE PROCEDURE update_oostats_on_insert();
CREATE TRIGGER update_oostats_delete AFTER DELETE ON oorders FOR EACH ROW EXECUTE PROCEDURE update_oostats_on_delete();
CREATE TRIGGER update_mkt_stats_insert AFTER INSERT ON market FOR EACH ROW EXECUTE PROCEDURE update_market_stats_on_insert();
CREATE TRIGGER update_mkt_stats_delete AFTER DELETE ON market FOR EACH ROW EXECUTE PROCEDURE update_market_stats_on_delete();
