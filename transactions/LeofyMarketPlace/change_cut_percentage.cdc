import LeofyMarketPlace from "../../contracts/LeofyMarketPlace.cdc"


transaction(cutPercentage: UFix64) {
    let adminRef: &LeofyMarketPlace.LeofyMarketPlaceAdmin

    prepare(signer: AuthAccount) {
        self.adminRef = signer.borrow<&LeofyMarketPlace.LeofyMarketPlaceAdmin>(from: LeofyMarketPlace.AdminStoragePath)
			?? panic("Could not borrow reference to the owner's AdminLeofyMarketPlace!")
    }
    execute {
        self.adminRef.changePercentage(cutPercentage)
    }
}