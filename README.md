# gpt-service

to pass context, use:
'''
{
    "input": "what gpt model are you 3.5 or 4?",
    "model": "gpt-4", // or "gpt-3.5" for gpt3dot5turbo - it's also a default one, or "gpt-4-23k
    "context": [
        {
            "isFromUser": true,
            "message": "My name is Maksim, in responses call me by my name"
        }
    ]
}
'''