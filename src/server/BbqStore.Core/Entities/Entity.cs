using System;

namespace BbqStore.Core.Entities
{
    public class Entity
    {
        public virtual Guid Id { get; set; }
        public virtual string CreatedBy { get; set; }
        public virtual DateTimeOffset CreatedDate { get; set; }
        public virtual string ModifiedBy { get; set; }
        public virtual DateTimeOffset ModifiedDate { get; set; }
        public virtual bool IsDeleted { get; set; }
    }
}