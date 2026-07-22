package main

func researchInstructions() string {
	return `Research the topic for a deeply technical "learning in public" blog post.

The writer wants to explain:
- what the technology is and why it is worth learning;
- how it works internally, including architecture, data flow, important abstractions, and design decisions;
- technical problems, limitations, failure modes, and trade-offs encountered when using it;
- surprising ideas, common misconceptions, and useful mental models discovered while learning.

Research requirements:
- Prefer official documentation, specifications, papers, source code, and original engineering posts.
- Explain mechanisms and causality, not only features or definitions.
- Include concrete technical examples where evidence supports them.
- Assign every source a stable ID such as S1.
- Every factual finding must reference one or more source IDs.
- Separate sourced facts from inference and unresolved questions.
- Do not invent the writer's personal experience. Instead, suggest learning angles or reflection prompts the writer can personalize.

Return only valid JSON with exactly this structure:
{
  "topic": "the researched topic",
  "summary": "a concise mental model and why the topic matters",
  "findings": [
    {
      "claim": "a detailed technical finding explaining what happens, how it works, and why it matters",
      "source_ids": ["S1"]
    }
  ],
  "conflicts": [
    "uncertain evidence, disagreement, limitation, or unresolved technical question"
  ],
  "angles": [
    "a learning insight, surprising discovery, misconception, or personal reflection prompt"
  ],
  "sources": [
    {
      "id": "S1",
      "title": "source title",
      "url": "https://example.com"
    }
  ]
}

Do not return Markdown fences or commentary outside the JSON.`
}
