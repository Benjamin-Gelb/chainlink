package models

// Yoinked from: https://github.com/langchain-ai/langchain/blob/master/libs/community/langchain_community/callbacks/openai_info.py

var MODEL_CHAT_VS_COMPLETION = map[string]flaot64{
	"gpt-4":                  0.03,
	"gpt-4-0314":             0.03,
	"gpt-4-0613":             0.03,
	"gpt-4-32k":              0.06,
	"gpt-4-32k-0314":         0.06,
	"gpt-4-32k-0613":         0.06,
	"gpt-4-vision-preview":   0.01,
	"gpt-4-1106-preview":     0.01,
	"gpt-4-0125-preview":     0.01,
	"gpt-4-turbo-preview":    0.01,
	"gpt-4-turbo":            0.01,
	"gpt-4-turbo-2024-04-09": 0.01,
}

var MODEL_COST_PER_1K_TOKENS = map[string]float64{
	// GPT-4 input
	"gpt-4":                  0.03,
	"gpt-4-0314":             0.03,
	"gpt-4-0613":             0.03,
	"gpt-4-32k":              0.06,
	"gpt-4-32k-0314":         0.06,
	"gpt-4-32k-0613":         0.06,
	"gpt-4-vision-preview":   0.01,
	"gpt-4-1106-preview":     0.01,
	"gpt-4-0125-preview":     0.01,
	"gpt-4-turbo-preview":    0.01,
	"gpt-4-turbo":            0.01,
	"gpt-4-turbo-2024-04-09": 0.01,
	// GPT-4 output
	"gpt-4-completion":                  0.06,
	"gpt-4-0314-completion":             0.06,
	"gpt-4-0613-completion":             0.06,
	"gpt-4-32k-completion":              0.12,
	"gpt-4-32k-0314-completion":         0.12,
	"gpt-4-32k-0613-completion":         0.12,
	"gpt-4-vision-preview-completion":   0.03,
	"gpt-4-1106-preview-completion":     0.03,
	"gpt-4-0125-preview-completion":     0.03,
	"gpt-4-turbo-preview-completion":    0.03,
	"gpt-4-turbo-completion":            0.03,
	"gpt-4-turbo-2024-04-09-completion": 0.03,
	// GPT-3.5 input
	// gpt-3.5-turbo points at gpt-3.5-turbo-0613 until Feb 16, 2024.
	// Switches to gpt-3.5-turbo-0125 after.
	"gpt-3.5-turbo":          0.0015,
	"gpt-3.5-turbo-0125":     0.0005,
	"gpt-3.5-turbo-0301":     0.0015,
	"gpt-3.5-turbo-0613":     0.0015,
	"gpt-3.5-turbo-1106":     0.001,
	"gpt-3.5-turbo-instruct": 0.0015,
	"gpt-3.5-turbo-16k":      0.003,
	"gpt-3.5-turbo-16k-0613": 0.003,
	// GPT-3.5 output
	// gpt-3.5-turbo points at gpt-3.5-turbo-0613 until Feb 16, 2024.
	// Switches to gpt-3.5-turbo-0125 after.
	"gpt-3.5-turbo-completion":          0.002,
	"gpt-3.5-turbo-0125-completion":     0.0015,
	"gpt-3.5-turbo-0301-completion":     0.002,
	"gpt-3.5-turbo-0613-completion":     0.002,
	"gpt-3.5-turbo-1106-completion":     0.002,
	"gpt-3.5-turbo-instruct-completion": 0.002,
	"gpt-3.5-turbo-16k-completion":      0.004,
	"gpt-3.5-turbo-16k-0613-completion": 0.004,
	// Azure GPT-35 input
	"gpt-35-turbo":          0.0015, // Azure OpenAI version of ChatGPT
	"gpt-35-turbo-0301":     0.0015, // Azure OpenAI version of ChatGPT
	"gpt-35-turbo-0613":     0.0015,
	"gpt-35-turbo-instruct": 0.0015,
	"gpt-35-turbo-16k":      0.003,
	"gpt-35-turbo-16k-0613": 0.003,
	// Azure GPT-35 output
	"gpt-35-turbo-completion":          0.002, // Azure OpenAI version of ChatGPT
	"gpt-35-turbo-0301-completion":     0.002, // Azure OpenAI version of ChatGPT
	"gpt-35-turbo-0613-completion":     0.002,
	"gpt-35-turbo-instruct-completion": 0.002,
	"gpt-35-turbo-16k-completion":      0.004,
	"gpt-35-turbo-16k-0613-completion": 0.004,
	// Others
	"text-ada-001":     0.0004,
	"ada":              0.0004,
	"text-babbage-001": 0.0005,
	"babbage":          0.0005,
	"text-curie-001":   0.002,
	"curie":            0.002,
	"text-davinci-003": 0.02,
	"text-davinci-002": 0.02,
	"code-davinci-002": 0.02,
	// Fine Tuned input
	"babbage-002-finetuned":        0.0016,
	"davinci-002-finetuned":        0.012,
	"gpt-3.5-turbo-0613-finetuned": 0.012,
	"gpt-3.5-turbo-1106-finetuned": 0.012,
	// Fine Tuned output
	"babbage-002-finetuned-completion":        0.0016,
	"davinci-002-finetuned-completion":        0.012,
	"gpt-3.5-turbo-0613-finetuned-completion": 0.016,
	"gpt-3.5-turbo-1106-finetuned-completion": 0.016,
	// Azure Fine Tuned input
	"babbage-002-azure-finetuned":       0.0004,
	"davinci-002-azure-finetuned":       0.002,
	"gpt-35-turbo-0613-azure-finetuned": 0.0015,
	// Azure Fine Tuned output
	"babbage-002-azure-finetuned-completion":       0.0004,
	"davinci-002-azure-finetuned-completion":       0.002,
	"gpt-35-turbo-0613-azure-finetuned-completion": 0.002,
	// Legacy fine-tuned models
	"ada-finetuned-legacy":     0.0016,
	"babbage-finetuned-legacy": 0.0024,
	"curie-finetuned-legacy":   0.012,
	"davinci-finetuned-legacy": 0.12,
}
