"""
Test the BM25 retriever directly
"""
from haystack import Document
from haystack.document_stores.in_memory import InMemoryDocumentStore
from haystack.components.retrievers.in_memory import InMemoryBM25Retriever

# Create test documents
docs = [
    Document(content="You can place a bid by calling the bid() function in the smart contract.", 
             meta={"file": "bidding.sol"}),
    Document(content="Prize distribution is handled automatically after each round ends.",
             meta={"file": "prize.sol"}),
    Document(content="The contract uses ERC20 tokens for bidding.",
             meta={"file": "token.sol"}),
]

# Create document store and add docs
doc_store = InMemoryDocumentStore()
doc_store.write_documents(docs)

print(f"Documents in store: {doc_store.count_documents()}")

# Create retriever
retriever = InMemoryBM25Retriever(document_store=doc_store)

# Test queries
test_queries = [
    "how to bid",
    "prize distribution",
    "what tokens"
]

for query in test_queries:
    print(f"\nQuery: {query}")
    result = retriever.run(query=query, top_k=2)
    docs_found = result.get("documents", [])
    print(f"Found {len(docs_found)} documents")
    for doc in docs_found:
        print(f"  - {doc.meta.get('file')}: {doc.content[:50]}...")
