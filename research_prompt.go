package main

func researchInstructions() string {
	return `
	Prepare a source-grounded research dossier for a technically rigorous educational blog post.

	Audience:
	- Technically curious readers who want to understand how and why the topic works.
	- Explain necessary terminology and prerequisites without oversimplifying.

	Research requirements:
	- Prefer primary sources: official docs, specifications, papers, original announcement of dev team, etc.
	- Assign every source a stable ID such as S1.
	- Cite one or more source IDs for every factual claim.
	- Separate sourced facts from analysis or inference.
	- Do not invent missing technical details; explicitly identify research gaps.
	- Compare conflicting evidence and explain possible reasons for disagreement.
	- Include mechanisms, architecture, workflows, trade-offs, limitations, failure modes, and practical implications.
	- Include concrete technical examples where supported by evidence.

	Return these sections:
	1. Scope and assumptions
	2. Executive summary
	3. Prerequisites and key terminology
	4. Technical deep dive
	5. Evidence and benchmarks
	6. Trade-offs, limitations, and failure modes
	7. Practical examples
	8. Conflicting or uncertain information
	9. Potential blog thesis and angles
	10. Suggested learning-oriented outline
	11. Open research questions
	12. Sources

	Write detailed markdown suitable as source material for a later article, not as a polished final blog post.
	`
}
