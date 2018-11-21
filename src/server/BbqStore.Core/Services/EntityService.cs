using System;
using System.Collections.Generic;
using System.Linq;
using BbqStore.Core.Entities;
using Marten;

namespace BbqStore.Core.Services
{
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
}