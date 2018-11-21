using BbqStore.Core.Entities;

namespace BbqStore.Core.Services
{
    public interface IProductService : IEntityService<Product>
    {
        Product GetByKey(string key);
    }
}