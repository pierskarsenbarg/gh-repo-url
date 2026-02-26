import * as pulumi from "@pulumi/pulumi";
import * as github from "@pulumi/github";

const repo = new github.Repository("repo", {
    allowForking: true,
    hasIssues: true,
    hasProjects: true,
    hasWiki: true,
    name: "gh-repo-url",
    securityAndAnalysis: {
        secretScanning: {
            status: "disabled",
        },
        secretScanningPushProtection: {
            status: "disabled",
        },
    },
    visibility: "public",
}, {
    protect: true,
});

const defaultBranch = new github.BranchDefault("main-branch", {
    branch: "main",
    repository: repo.name
}, {
    protect: true
});