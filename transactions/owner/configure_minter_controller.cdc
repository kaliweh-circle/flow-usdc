// Masterminter uses this to configure which minter the minter controller manages

import FiatToken from 0x{{.FiatToken}}
import FiatTokenInterface from 0x{{.FiatTokenInterface}}

transaction (minter: UInt64, minterController: UInt64) {
    prepare(masterMinter: AuthAccount) {
        let mm = masterMinter.borrow<&FiatToken.MasterMinter{FiatTokenInterface.MasterMinter}>(from: FiatToken.MasterMinterStoragePath) 
            ?? panic ("no masterminter resource avaialble");

        mm.configureMinterController(minter: minter, minterController: minterController);
    }
    post {
        FiatToken.managedMinters[minterController] == minter : "minterController not configured"
    }
}