using System;
using System.Collections.Generic;
using BbqStore.Core.Entities;

namespace BbqStore.Core.Services
{
    public interface IEntityService<T> where T : Entity
    {
        T Save(T entity);
        void Delete(T entity);
        T GetById(Guid id);
        IEnumerable<T> GetAll();
    }
}