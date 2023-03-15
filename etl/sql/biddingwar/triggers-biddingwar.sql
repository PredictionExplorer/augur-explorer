CREATE TRIGGER bw_bid_insert AFTER INSERT ON bw_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_insert();
CREATE TRIGGER bw_bid_delete AFTER DELETE ON bw_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_delete();
CREATE TRIGGER bw_prize_claim_insert AFTER INSERT ON bw_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_insert();
CREATE TRIGGER bw_prize_claim_delete AFTER DELETE ON bw_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_delete();
CREATE TRIGGER bw_donation_received_insert AFTER INSERT ON bw_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_insert();
CREATE TRIGGER bw_donation_received_delete AFTER DELETE ON bw_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_delete();
