using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using BbqStore.Core.Entities;
using Marten;
using Microsoft.Extensions.Configuration;

namespace BbqStore.Core.Services
{
    public interface IConfigurationService
    {
        string Get(string key);
    }

    public class ConfigurationService : IConfigurationService
    {
        public ConfigurationService(IConfiguration config)
        {
            Configuration = config;
        }

        protected IConfiguration Configuration { get; set; }

        public string Get(string key)
        {
            return Configuration[key];
        }
    }

    public interface IEntityService<T> where T : Entity
    {
        T Save(T entity);
        void Delete(T entity);
        T GetById(Guid id);
        IEnumerable<T> GetAll();
    }

    public class EntityService<T> : IEntityService<T> where T : Entity
    {
        protected IDocumentSession DocumentSession { get; set; }

        public EntityService(IDocumentSession documentSession)
        {
            DocumentSession = documentSession;
        }

        public virtual T Save(T entity)
        {
            if(Guid.Empty == entity.Id)
            {
                entity.Id = Guid.NewGuid();
                entity.CreatedBy = "chef";
                entity.CreatedDate = DateTimeOffset.Now;

                if (entity is NamedEntity)
                {
                    var namedEntity = entity as NamedEntity;
                    if (String.IsNullOrEmpty(namedEntity.Key))
                        namedEntity.Key = namedEntity.Name.Replace(" ", "-").ToLower();
                }
            }

            entity.ModifiedBy = "chef";
            entity.ModifiedDate = DateTimeOffset.Now;

            DocumentSession.Store(entity);
            DocumentSession.SaveChanges();

            return entity;
        }

        public virtual void Delete(T entity)
        {
            entity.IsDeleted = true;
            Save(entity);
        }

        public virtual T GetById(Guid id)
        {
            return DocumentSession.Query<T>().FirstOrDefault(x => x.Id == id);
        }

        public virtual IEnumerable<T> GetAll()
        {
            return DocumentSession.Query<T>().ToList();
        }
    }

    public interface IProductService : IEntityService<Product>
    {
        Product GetByKey(string key);
    }

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
