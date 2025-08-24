import argparse
import os

from dotenv import load_dotenv
from google import genai
from google.genai import types


class RequestInstance:
    def __init__(self, contents: list, client: genai.Client, model: str):
        self.contents = contents
        self.client = client
        self.model = model


class ResponseInstance:
    def __init__(self, response: str, response_tokens: int, request_tokens: int):
        self.response = response
        self.response_tokens = response_tokens
        self.request_tokens = request_tokens


def get_reponse(request: RequestInstance) -> ResponseInstance:
    response_model = request.client.models.generate_content(
        model=request.model, contents=request.contents
    )
    response_text = response_model.text
    prompt_tokens = response_model.usage_metadata.prompt_token_count
    response_tokens = response_model.usage_metadata.candidates_token_count

    return ResponseInstance(response_text, response_tokens, prompt_tokens)


def log_response(response: ResponseInstance):
    print(f"The model responded:\n{response.response}")
    # print(f"The response consumed:\n{response.response_tokens} Tokens")
    # print(f"The prompt consumed:\n{response.request_tokens} Tokens")
    if args.verbose:
        print(f"User prompt: {args.prompt}")
        print(f"Prompt tokens: {response.request_tokens}")
        print(f"Response tokens: {response.response_tokens}")


def main():
    load_dotenv()
    api_key = os.environ.get("GEMINI_API_KEY")
    model = os.environ.get("MODEL")
    client = genai.Client(api_key=api_key)
    messages = [types.Content(role="user", parts=[types.Part(text=args.prompt)])]
    request = RequestInstance(contents=messages, model=model, client=client)
    response = get_reponse(request)
    log_response(response)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="CLI Tool Used to Create Project README.md Files Using AI"
    )
    parser.add_argument("prompt", help="The user prompt sent to Gemini.")
    parser.add_argument(
        "--verbose", action="store_true", help="Enable verbose logging."
    )
    args = parser.parse_args()
    main()
