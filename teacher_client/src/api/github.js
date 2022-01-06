import {Octokit} from '@octokit/core'

const octokit = new Octokit()

export function getRepoReleases() {
  return octokit.request('GET /repos/{owner}/{repo}/releases', {
    owner: 'gonearewe',
    repo: 'EasyTesting'
  })
}
