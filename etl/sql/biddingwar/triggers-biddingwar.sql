CREATE TRIGGER bw_bid_insert AFTER INSERT ON bw_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_insert();
CREATE TRIGGER bw_bid_delete AFTER DELETE ON bw_bid FOR EACH ROW EXECUTE PROCEDURE on_bid_delete();
CREATE TRIGGER bw_prize_claim_insert AFTER INSERT ON bw_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_insert();
CREATE TRIGGER bw_prize_claim_delete AFTER DELETE ON bw_prize_claim FOR EACH ROW EXECUTE PROCEDURE on_prize_claim_delete();
CREATE TRIGGER bw_donation_received_insert AFTER INSERT ON bw_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_insert();
CREATE TRIGGER bw_donation_received_delete AFTER DELETE ON bw_donation_received FOR EACH ROW EXECUTE PROCEDURE on_donation_received_delete();
CREATE TRIGGER bw_nft_donation_insert AFTER INSERT ON bw_nft_donation FOR EACH ROW EXECUTE PROCEDURE on_nft_donation_insert();
CREATE TRIGGER bw_nft_donation_delete AFTER DELETE ON bw_nft_donation FOR EACH ROW EXECUTE PROCEDURE on_nft_donation_delete();
CREATE TRIGGER bw_raffle_deposit_insert AFTER INSERT ON bw_raffle_deposit FOR EACH ROW EXECUTE PROCEDURE on_raffle_deposit_insert();
CREATE TRIGGER bw_raffle_deposit_delete AFTER DELETE ON bw_raffle_deposit FOR EACH ROW EXECUTE PROCEDURE on_raffle_deposit_delete();
CREATE TRIGGER bw_raffle_nft_winner_insert AFTER INSERT ON bw_raffle_nft_winner FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_winner_insert();
CREATE TRIGGER bw_raffle_nft_winner_delete AFTER DELETE ON bw_raffle_nft_winner FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_winner_delete();
CREATE TRIGGER bw_raffle_nft_claimed_insert AFTER INSERT ON bw_raffle_nft_claimed FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_claimed_insert();
CREATE TRIGGER bw_raffle_nft_claimed_delete AFTER DELETE ON bw_raffle_nft_claimed FOR EACH ROW EXECUTE PROCEDURE on_raffle_nft_claimed_delete();
CREATE TRIGGER bw_transer_insert AFTER INSERT ON bw_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc721transfer_insert();
CREATE TRIGGER bw_transfer_delete AFTER DELETE ON bw_transfer FOR EACH ROW EXECUTE PROCEDURE on_erc721transfer_delete();
