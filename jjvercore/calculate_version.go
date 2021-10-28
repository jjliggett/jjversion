package jjvercore

var oss osInterface
var gs gitServiceInterface

func init() {
	oss = osService{}
	gs = gitService{}
}

func CalculateVersion() VersionOutput {
	path, err := oss.getwd()
	CheckIfError(err)

	vs := getVersionSettings()

	r, err := gs.openRepository(path)
	CheckIfError(err)

	ref, err := gs.getRepositoryHead(r)
	CheckIfError(err)

	if ref.Name().IsBranch() {
		branchVersion, success := GetVersionFromReleaseBranch(ref.Name().Short())
		if success {
			return VersionOutput{branchVersion.Major, branchVersion.Minor, branchVersion.Patch, ref.Hash().String()}
		}
	}

	tagrefs, err := gs.getRepositoryTags(r)
	CheckIfError(err)

	tags, err := gs.getRepositoryTagObjects(r)
	CheckIfError(err)

	versions := getVersionTags(tagrefs, tags)

	commit_version, ok := versions[ref.Hash().String()]
	if ok {
		return VersionOutput{commit_version.Major, commit_version.Minor, commit_version.Patch, ref.Hash().String()}
	}

	cIter, err := gs.getRepositoryCommits(r)
	CheckIfError(err)

	v := version{0, 0, 0}

	v = incrementVersionByCommitMessages(cIter, versions, vs, v)

	return VersionOutput{v.Major, v.Minor, v.Patch, ref.Hash().String()}
}
