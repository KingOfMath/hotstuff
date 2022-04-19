package twophasehotstuff

import (
	"github.com/relab/hotstuff/consensus"
	"github.com/relab/hotstuff/modules"
)

func init() {
	modules.RegisterModule("twophasehotstuff", New)
}

type TwoPhaseHotStuff struct {
	mods *consensus.Modules

	bLock *consensus.Block
}

// New returns a new FastHotStuff instance.
func New() consensus.Rules {
	return &TwoPhaseHotStuff{
		bLock: consensus.GetGenesis(),
	}
}

func (t *TwoPhaseHotStuff) InitConsensusModule(mods *consensus.Modules, _ *consensus.OptionsBuilder) {
	t.mods = mods
}

func (t *TwoPhaseHotStuff) VoteRule(proposal consensus.ProposeMsg) bool {
	block := proposal.Block

	if block.View() < t.mods.Synchronizer().View() {
		t.mods.Logger().Info("VoteRule: block view too low")
		return false
	}

	parent, ok := t.mods.BlockChain().Get(block.QuorumCert().BlockHash())
	// should extend the previous highest QC
	return ok && parent.View() >= t.bLock.View() && t.mods.BlockChain().Extends(block, t.bLock)
}

func (t *TwoPhaseHotStuff) CommitRule(block *consensus.Block) *consensus.Block {
	// parent
	parent, ok := t.mods.BlockChain().Get(block.QuorumCert().BlockHash())
	if !ok {
		return nil
	}
	if parent.View() > t.bLock.View() {
		t.bLock = parent
		t.mods.Logger().Debug("Prepare-Locked: ", parent)
	}

	// grandparent
	grandparent, ok := t.mods.BlockChain().Get(parent.QuorumCert().BlockHash())
	if !ok {
		return nil
	}
	if grandparent.View()+1 == parent.View() {
		t.mods.Logger().Debug("Commit: ", grandparent)
		return grandparent
	}
	return nil
}

func (t *TwoPhaseHotStuff) ChainLength() int {
	return 2
}
