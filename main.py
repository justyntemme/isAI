import argparse
import openai

def is_ai_generated(api_key, paragraph):
    openai.api_key = api_key
    prompt = "Is the following paragraph human-written or AI-generated?\n\n" + paragraph + "\n\nAnswer: "
    response = openai.Completion.create(
        engine="text-davinci-003",
        prompt=prompt,
        temperature=2,
        max_tokens=40,
        n=1,
        stop=None,
        timeout=10
    )
    answer = response.choices[0].text.strip().lower()
    if answer == "ai-generated":
        return True
    elif answer == "human-written":
        return False
    else:
        return None

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--key", type=str, required=True, help="OpenAI API key")
    parser.add_argument("--text", type=str, required=True, help="Paragraph text to analyze")
    args = parser.parse_args()

    if is_ai_generated(args.key, args.text):
        print("The paragraph is AI-generated.")
    else:
        print("The paragraph is human-written.")
