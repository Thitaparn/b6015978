// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/playlist"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/resolution"
	"github.com/tanapon395/playlist-video/ent/video"
)

// PlaylistVideoCreate is the builder for creating a Playlist_Video entity.
type PlaylistVideoCreate struct {
	config
	mutation *PlaylistVideoMutation
	hooks    []Hook
}

// SetAddedTime sets the added_time field.
func (pvc *PlaylistVideoCreate) SetAddedTime(t time.Time) *PlaylistVideoCreate {
	pvc.mutation.SetAddedTime(t)
	return pvc
}

// SetPlaylistID sets the playlist edge to Playlist by id.
func (pvc *PlaylistVideoCreate) SetPlaylistID(id int) *PlaylistVideoCreate {
	pvc.mutation.SetPlaylistID(id)
	return pvc
}

// SetNillablePlaylistID sets the playlist edge to Playlist by id if the given value is not nil.
func (pvc *PlaylistVideoCreate) SetNillablePlaylistID(id *int) *PlaylistVideoCreate {
	if id != nil {
		pvc = pvc.SetPlaylistID(*id)
	}
	return pvc
}

// SetPlaylist sets the playlist edge to Playlist.
func (pvc *PlaylistVideoCreate) SetPlaylist(p *Playlist) *PlaylistVideoCreate {
	return pvc.SetPlaylistID(p.ID)
}

// SetVideoID sets the video edge to Video by id.
func (pvc *PlaylistVideoCreate) SetVideoID(id int) *PlaylistVideoCreate {
	pvc.mutation.SetVideoID(id)
	return pvc
}

// SetNillableVideoID sets the video edge to Video by id if the given value is not nil.
func (pvc *PlaylistVideoCreate) SetNillableVideoID(id *int) *PlaylistVideoCreate {
	if id != nil {
		pvc = pvc.SetVideoID(*id)
	}
	return pvc
}

// SetVideo sets the video edge to Video.
func (pvc *PlaylistVideoCreate) SetVideo(v *Video) *PlaylistVideoCreate {
	return pvc.SetVideoID(v.ID)
}

// SetResolutionID sets the resolution edge to Resolution by id.
func (pvc *PlaylistVideoCreate) SetResolutionID(id int) *PlaylistVideoCreate {
	pvc.mutation.SetResolutionID(id)
	return pvc
}

// SetNillableResolutionID sets the resolution edge to Resolution by id if the given value is not nil.
func (pvc *PlaylistVideoCreate) SetNillableResolutionID(id *int) *PlaylistVideoCreate {
	if id != nil {
		pvc = pvc.SetResolutionID(*id)
	}
	return pvc
}

// SetResolution sets the resolution edge to Resolution.
func (pvc *PlaylistVideoCreate) SetResolution(r *Resolution) *PlaylistVideoCreate {
	return pvc.SetResolutionID(r.ID)
}

// Mutation returns the PlaylistVideoMutation object of the builder.
func (pvc *PlaylistVideoCreate) Mutation() *PlaylistVideoMutation {
	return pvc.mutation
}

// Save creates the Playlist_Video in the database.
func (pvc *PlaylistVideoCreate) Save(ctx context.Context) (*Playlist_Video, error) {
	if err := pvc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Playlist_Video
	)
	if len(pvc.hooks) == 0 {
		node, err = pvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvc.mutation = mutation
			node, err = pvc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pvc.hooks) - 1; i >= 0; i-- {
			mut = pvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pvc *PlaylistVideoCreate) SaveX(ctx context.Context) *Playlist_Video {
	v, err := pvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvc *PlaylistVideoCreate) preSave() error {
	if _, ok := pvc.mutation.AddedTime(); !ok {
		return &ValidationError{Name: "added_time", err: errors.New("ent: missing required field \"added_time\"")}
	}
	return nil
}

func (pvc *PlaylistVideoCreate) sqlSave(ctx context.Context) (*Playlist_Video, error) {
	pv, _spec := pvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pvc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pv.ID = int(id)
	return pv, nil
}

func (pvc *PlaylistVideoCreate) createSpec() (*Playlist_Video, *sqlgraph.CreateSpec) {
	var (
		pv    = &Playlist_Video{config: pvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playlist_video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist_video.FieldID,
			},
		}
	)
	if value, ok := pvc.mutation.AddedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: playlist_video.FieldAddedTime,
		})
		pv.AddedTime = value
	}
	if nodes := pvc.mutation.PlaylistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist_video.PlaylistTable,
			Columns: []string{playlist_video.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist_video.VideoTable,
			Columns: []string{playlist_video.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.ResolutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist_video.ResolutionTable,
			Columns: []string{playlist_video.ResolutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resolution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pv, _spec
}

// PlaylistVideoCreateBulk is the builder for creating a bulk of Playlist_Video entities.
type PlaylistVideoCreateBulk struct {
	config
	builders []*PlaylistVideoCreate
}

// Save creates the Playlist_Video entities in the database.
func (pvcb *PlaylistVideoCreateBulk) Save(ctx context.Context) ([]*Playlist_Video, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pvcb.builders))
	nodes := make([]*Playlist_Video, len(pvcb.builders))
	mutators := make([]Mutator, len(pvcb.builders))
	for i := range pvcb.builders {
		func(i int, root context.Context) {
			builder := pvcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*PlaylistVideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pvcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pvcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pvcb *PlaylistVideoCreateBulk) SaveX(ctx context.Context) []*Playlist_Video {
	v, err := pvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}