"""
Debug script to test indexing and querying
"""
import asyncio
import logging
from rag_pipeline import CosmicRAGPipeline

# Set up logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)

async def main():
    print("=" * 60)
    print("Cosmic FAQ Bot - Debug Script")
    print("=" * 60)
    print()
    
    # Initialize pipeline
    print("1. Initializing RAG pipeline...")
    pipeline = CosmicRAGPipeline()
    print(f"   ✓ Pipeline initialized")
    print()
    
    # Check if indexed
    doc_count = pipeline.document_store.count_documents()
    print(f"2. Current document count: {doc_count}")
    print(f"   Is indexed: {pipeline.is_indexed()}")
    print()
    
    # Index if needed
    if not pipeline.is_indexed():
        print("3. No documents found. Starting indexing...")
        print("   This may take a few minutes...")
        print()
        await pipeline.index_repositories()
        doc_count = pipeline.document_store.count_documents()
        print(f"   ✓ Indexing complete! Total documents: {doc_count}")
    else:
        print("3. Documents already indexed")
    print()
    
    # Show sample documents
    if doc_count > 0:
        print("4. Sample documents:")
        all_docs = pipeline.document_store.filter_documents()
        for i, doc in enumerate(all_docs[:5], 1):
            print(f"   {i}. {doc.meta.get('repository')}/{doc.meta.get('file_path')}")
            print(f"      Type: {doc.meta.get('file_type')}, Length: {len(doc.content)} chars")
    else:
        print("4. ⚠ WARNING: No documents were indexed!")
        print("   Possible issues:")
        print("   - Git clone failed (check internet connection)")
        print("   - No supported files found in repositories")
        print("   - File filtering too strict")
        return
    print()
    
    # Test queries
    print("5. Testing sample queries:")
    test_questions = [
        "How can I bid?",
        "smart contract",
        "prize distribution",
        "What is this project about?"
    ]
    
    for question in test_questions:
        print(f"\n   Q: {question}")
        result = await pipeline.query(question)
        print(f"   Retrieved sources: {len(result.get('sources', []))}")
        if result.get('sources'):
            print(f"   First source: {result['sources'][0]}")
        answer_preview = result['answer'][:150] + "..." if len(result['answer']) > 150 else result['answer']
        print(f"   Answer preview: {answer_preview}")
    
    print()
    print("=" * 60)
    print("Debug complete!")
    print("=" * 60)

if __name__ == "__main__":
    asyncio.run(main())
