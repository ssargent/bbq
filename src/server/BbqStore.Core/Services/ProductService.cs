using System.Linq;
using BbqStore.Core.Entities;
using Marten;

namespace BbqStore.Core.Services
{
    public class ProductService : EntityService<Product>, IProductService
    {
        public ProductService(IDocumentSession documentSession) : base(documentSession)
        {
        }

        public Product GetByKey(string key)
        {
            return DocumentSession.Query<Product>().FirstOrDefault(x => x.Key == key);
        }
    }
}